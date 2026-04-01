/**
 * sleeping_barbers.go demonstrates a solution to the "Sleeping Barbers" problem,
 * formulated by Edsger Dijkstra:
 * https://en.wikipedia.org/wiki/Sleeping_barber_problem
 */

package misc

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// barbers is the main method for the Sleeping Barbers problem.
// Number of barbers, number of waiting chairs, number of customers, and
// wait time must be supplied by the caller.
func barbers(fNumBarbers, fNumChairs, fNumCustomers, fWaitTime int) error {

	shop := newShop(fNumBarbers, fWaitTime, fNumChairs)
	var wg sync.WaitGroup

	for j := 0; j < fNumBarbers; j++ {
		go func(id int, shop *barberShop) {
			barber(id, shop)
		}(j, &shop)
	}

	for i := 0; i < fNumCustomers; i++ {
		rng := rand.Intn(1000)
		d := time.Duration(rng)
		time.Sleep(d)
		wg.Add(1)
		go func(id int, shop *barberShop) {
			customer(id, shop)
			fmt.Printf("customer [%d]: is happily on their way home\n", id)
			wg.Done()
		}(i, &shop)
	}

	wg.Wait()
	fmt.Println("Wait group off")
	shop.quit()

	fmt.Printf("Number of lost customers: %d\n", shop.dropOff)

	return nil
}

type queue struct {
	nums []int
}

func (q *queue) Size() int {
	return len(q.nums)
}

func (q *queue) Enqueue(num int) {
	for i := 0; i < len(q.nums); i++ {
		if num == q.nums[i] {
			return
		}
	}

	q.nums = append(q.nums, num)
}

func (q *queue) Dequeue() int {
	if q.Size() == 0 {
		return -1
	}
	num := q.nums[0]
	q.nums = q.nums[1:]
	return num
}

type barberShop struct {
	waitTime             time.Duration
	maxWaitingChairs     int
	numBarbers           int
	waitingCustomers     queue
	availableBarbers     queue
	barberChairs         []int
	isChairActive        []bool
	isMoneyPaid          []bool
	lock                 *sync.Mutex
	condCustomersWaiting *sync.Cond
	condCustomerServed   []*sync.Cond
	condBarberSleeping   []*sync.Cond
	condBarberPaid       []*sync.Cond
	dropOff              int
	isQuittingTime       bool
}

func newShop(numBarbers, waitTime, waitingChairs int) barberShop {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	var d time.Duration = time.Duration(waitTime * int(time.Millisecond))
	thisShop := barberShop{
		numBarbers:           numBarbers,
		maxWaitingChairs:     waitingChairs,
		waitTime:             d,
		lock:                 &mu,
		condCustomersWaiting: cond,
		isQuittingTime:       false,
	}

	thisShop.condCustomerServed = make([]*sync.Cond, numBarbers)
	thisShop.condBarberSleeping = make([]*sync.Cond, numBarbers)
	thisShop.condBarberPaid = make([]*sync.Cond, numBarbers)
	thisShop.isChairActive = make([]bool, numBarbers)
	thisShop.isMoneyPaid = make([]bool, numBarbers)
	thisShop.barberChairs = make([]int, numBarbers)

	for i := 0; i < numBarbers; i++ {
		thisShop.barberChairs[i] = -1
		thisShop.condBarberPaid[i] = sync.NewCond(thisShop.lock)
		thisShop.condBarberSleeping[i] = sync.NewCond(thisShop.lock)
		thisShop.condCustomerServed[i] = sync.NewCond(thisShop.lock)
		thisShop.availableBarbers.Enqueue(i)
	}

	return thisShop
}

func (s *barberShop) quit() {
	s.lock.Lock()
	fmt.Println("Quitting time!")
	s.isQuittingTime = true
	for i := 0; i < s.numBarbers; i++ {
		s.condBarberSleeping[i].Signal()
	}
	s.lock.Unlock()
}

func (s *barberShop) areAnyChairsEmpty() bool {
	for i := 0; i < len(s.barberChairs); i++ {
		if s.barberChairs[i] == -1 {
			return true
		}
	}
	return false
}

func (s *barberShop) visitShop(id int) int {
	s.lock.Lock()

	if s.waitingCustomers.Size() >= s.maxWaitingChairs {
		fmt.Printf("customer [%d]: leaves the shop because of no available waiting chairs\n", id)
		s.dropOff++
		s.lock.Unlock()
		return -1
	}

	if !s.areAnyChairsEmpty() || s.waitingCustomers.Size() > 0 {
		s.waitingCustomers.Enqueue(id)
		fmt.Printf("customer [%d]: takes a waiting chair.\n", id)
		fmt.Println("# waiting seats available: ",
			s.maxWaitingChairs-s.waitingCustomers.Size())
		s.condCustomersWaiting.Wait()
		s.waitingCustomers.Dequeue()
	}

	barberID := s.availableBarbers.Dequeue()
	fmt.Printf("customer [%d]: moves to [%d] service chair\n", id, barberID)
	s.barberChairs[barberID] = id
	s.isChairActive[barberID] = true
	s.isMoneyPaid[barberID] = false

	s.condBarberSleeping[barberID].Signal()

	s.lock.Unlock()

	return barberID
}

func (s *barberShop) leaveShop(id int, barberID int) {
	s.lock.Lock()

	fmt.Printf("customer [%d]: waits for barber[%d] to be done with haircut\n",
		id, barberID)

	for {
		if !s.isChairActive[barberID] {
			break
		}
		s.condCustomerServed[barberID].Wait()
	}

	s.isMoneyPaid[barberID] = true

	fmt.Printf("customer [%d]: says good-bye to barber [%d]\n",
		id, barberID)

	s.condBarberPaid[barberID].Signal()
	s.lock.Unlock()
}

func (s *barberShop) helloCustomer(id int) {
	s.lock.Lock()

	if s.waitingCustomers.Size() == 0 && s.barberChairs[id] < 0 {
		fmt.Printf("barber  [%d]: sleeps because of no customers.\n", id)
		s.condBarberSleeping[id].Wait()
	}

	if s.isQuittingTime {
		return
	}

	if s.barberChairs[id] < 0 {
		s.condBarberSleeping[id].Wait()
	}

	customerID := s.barberChairs[id]
	fmt.Printf("barber  [%d]: starts haircut for [%d]\n", id, customerID)
	s.lock.Unlock()
}

func (s *barberShop) byeCustomer(id int) {
	s.lock.Lock()

	customerID := s.barberChairs[id]
	s.isChairActive[id] = false
	fmt.Printf("barber  [%d]: says he's done with haircut service for [%d]\n",
		id, customerID)

	s.condCustomerServed[id].Signal()

	for {
		if s.isMoneyPaid[id] {
			break
		}
		s.condBarberPaid[id].Wait()
	}

	s.barberChairs[id] = -1
	s.isChairActive[id] = false
	s.isMoneyPaid[id] = false
	fmt.Printf("barber  [%d]: calls in another customer\n", id)

	s.availableBarbers.Enqueue(id)

	s.condCustomersWaiting.Signal()
	s.lock.Unlock()
}

func customer(id int, shop *barberShop) {
	barberID := shop.visitShop(id)

	if barberID >= 0 {
		shop.leaveShop(id, barberID)
	}
}

func barber(id int, shop *barberShop) {
	for {
		shop.helloCustomer(id)
		time.Sleep(shop.waitTime)

		if shop.isQuittingTime {
			break
		}

		shop.byeCustomer(id)

		if shop.isQuittingTime {
			break
		}
	}
}
