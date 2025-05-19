async function myFunction() {
  try {
    let response = await fetch('http://127.0.0.1:5000');
    let data = await response.json();
    return data; // Returns fetched data
  } catch (error) {
    console.error('Error:', error);
  }
}

async function button() {
  let result = await myFunction();  // Wait for the data to be fetched
  console.log(result.fstop);  // Now prints the actual response
}

document.getElementById("call").addEventListener("click", button);
