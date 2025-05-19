Kameramodell = document.getElementById("Kameramodell")
Verwendetes_Objektiv = document.getElementById("Verwendetes Objektiv")
Belichtungszeit = document.getElementById("Belichtungszeit")
Blende = document.getElementById("Blende")
ISO = document.getElementById("ISO")
Beleuchtung = document.getElementById("Beleuchtung")
Fokuspunkt = document.getElementById("Fokuspunkt")
Fotogenre = document.getElementById("Fotogenre")
save = document.getElementById("save")

function handle_save() {
    Kameramodell = document.getElementById("Kameramodell").value
Verwendetes_Objektiv = document.getElementById("Verwendetes Objektiv").value
Belichtungszeit = document.getElementById("Belichtungszeit").value
Blende = document.getElementById("Blende").value
ISO = document.getElementById("ISO").value
Beleuchtung = document.getElementById("Beleuchtung").value
Fokuspunkt = document.getElementById("Fokuspunkt").value
Fotogenre = document.getElementById("Fotogenre").value
const data = {
  Kameramodell,
  Verwendetes_Objektiv,
  Belichtungszeit,
  Blende,
  ISO,
  Beleuchtung,
  Fokuspunkt,
  Fotogenre
};

fetch("http://127.0.0.1:5000/add_image", {
  method: "POST",
  headers: {
    "Content-Type": "application/json"
  },
  body: JSON.stringify(data)
})
.then(response => response.json())
.then(result => {
  console.log("Success:", result);
})
.catch(error => {
  console.error("Error:", error);
});

}

save.addEventListener("click", handle_save)