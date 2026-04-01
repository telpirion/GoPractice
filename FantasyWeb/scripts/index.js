"use strict";

window.onload = function () {
    console.log("Welcome to Fantasy Maps!")

    var ui = new firebaseui.auth.AuthUI(firebase.auth());
    var uiConfig = {
        callbacks: {
          signInSuccessWithAuthResult: function(authResult, redirectUrl) {
            // User successfully signed in.
            // Return type determines whether we continue the redirect automatically
            // or whether we leave that to developer to handle.
            return true;
          },
          uiShown: function() {
            document.getElementById('loader').style.display = 'none';
          }
        },
        signInFlow: 'popup',
        signInSuccessUrl: '/submit-map',
        signInOptions: [
          firebase.auth.GoogleAuthProvider.PROVIDER_ID,
        ],
        tosUrl: '/terms-of-service',
        privacyPolicyUrl: '/privacy-policy'
      };


      ui.start('#firebaseui-auth-container', uiConfig);
}