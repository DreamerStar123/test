const admin = require('firebase-admin');
const serviceAccount = require('./turbo-tabs-firebase-adminsdk-vk1px-e1798541d6.json'); // Path to your service account JSON file

admin.initializeApp({
  credential: admin.credential.cert(serviceAccount),
});

admin.auth().listUsers()
    .then((listUsersResult) => {
        listUsersResult.users.forEach((userRecord) => {
            console.log("User:", userRecord.toJSON());
        });
    })
    .catch((error) => {
        console.error("Error listing users:", error);
    });
