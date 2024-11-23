// URL of the API endpoint
const url = 'https://userInfo-hhjsyj7q4q-uc.a.run.app';

// Fetch data from the API
fetch(url)
  .then(response => {
    // Check if the response is OK (status in the range 200-299)
    return response.json();
  })
  .then(data => {
    // Handle the parsed JSON data
    console.log(data);
  })
  .catch(error => {
    // Handle any errors that occurred during the fetch
    console.error('There was a problem with the fetch operation:', error);
  });