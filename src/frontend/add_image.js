save = document.getElementById("save").addEventListener("click", handle_save)
function handle_save() {
  const Kameramodell = document.getElementById("Kameramodell").value
  const Verwendetes_Objektiv = document.getElementById("Verwendetes Objektiv").value
  const Belichtungszeit = document.getElementById("Belichtungszeit").value
  const Blende = document.getElementById("Blende").value
  const ISO = document.getElementById("ISO").value
  const Beleuchtung = document.getElementById("Beleuchtung").value
  const Fokuspunkt = document.getElementById("Fokuspunkt").value
  const Fotogenre = document.getElementById("Fotogenre").value
  const image_address = document.getElementById("image_location").value
  const data = {
    image_address,
    Kameramodell,
    Verwendetes_Objektiv,
    Belichtungszeit,
    Blende,
    ISO,
    Beleuchtung,
    Fokuspunkt,
    Fotogenre
  };
  //write to the server
  fetch("http://127.0.0.1:5000/add_image", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  })
    .then(response => response.json())
    .then(result => {
      console.log("Success:", result.id);
      window.open("http://localhost:2020/frontend/display_image.html?content_id=" + result.id, "_blank");
    })
    .catch(error => {
      console.error("Error:", error);
    });

}