console.log("triggerd");
const url = window.location.href;
async function myFunction(id) {
  try {
    let response = await fetch('http://127.0.0.1:5000/' + id);
    let data = await response.json();
    console.log(data);
    return data; // Returns fetched data
  } catch (error) {
    console.error('Error:', error);
  }
}

const urlParams = new URLSearchParams(window.location.search);
const contentId = urlParams.get('content_id');

async function run() {
  let result = await myFunction(contentId);
  //console.log(result)
  document.getElementById('image').src = result.location;
  document.getElementById('textbox2').textContent =
    "Kameramodell: " + result.info.ISO
    + "\nVerwendetes Objektiv: " + result.info.camera_model
    + "\nBelichtungszeit: " + result.info.lens_used
    + "\nBlende: " + result.info.shutter_speed
    + "\nISO: " + result.info.aperture
    + "\nBeleuchtung: " + result.info.lighting
    + "\nFokuspunkt: " + result.info.focal_point
    + "\nFotogenre: " + result.info.photo_genre
}
run();
function info(text) {
  const info = document.getElementById('info');
  document.getElementById('toast').classList.remove('hidden'); // Show the alert
  info.querySelector('span').textContent = text;
  setTimeout(() => {
    document.getElementById('toast').classList.add("hidden");
  }, 1000);
}
function eventy() {
  navigator.clipboard.writeText(url)
    .then(() => {
      info("Addresse kopiert")
    })
    .catch(err => {
      console.error("Failed to copy: ", err);
    });


}
document.getElementById("share").addEventListener("click", eventy)











