import { initializeApp } from 'firebase/app';
import { getAuth, signInWithPopup, OAuthProvider, GoogleAuthProvider } from 'firebase/auth';

const firebaseConfig = {
  apiKey: "AIzaSyBZ1ET16vIOtnOMws9XvvGoANS2PGfTt_M",
  authDomain: "turbo-tabs.firebaseapp.com",
  projectId: "turbo-tabs",
  storageBucket: "turbo-tabs.appspot.com",
  messagingSenderId: "727774901928",
  appId: "1:727774901928:web:4e34f0d692b9fb314f349d",
  measurementId: "G-1WEF9BKD7X"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const auth = getAuth(app);

// Microsoft Sign-In
// const provider = new OAuthProvider('microsoft.com');
const provider = new GoogleAuthProvider();

signInWithPopup(auth, provider)
    .then((result) => {
        const credential = GoogleAuthProvider.credentialFromResult(result);
        console.log('credential ===================>', credential);
        console.log('result.user ===================>', result.user); // Signed-in user info
    })
    .catch((error) => {
        console.error(error); // Handle errors
    });