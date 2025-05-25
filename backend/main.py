import uuid
from flask import Flask, jsonify, request
from flask_cors import CORS # type: ignore
from image import image, user
app = Flask(__name__)
_ = CORS(app)
images = {"example": image("example", "	http://0.0.0.0:2020/images/866-536x354.jpg", {"fstop": True}, {
    "camera_model": "Canon EOS R5",
    "lens_used": "RF 50mm f/1.2L USM",
    "shutter_speed": "1/250 sec",
    "aperture": "f/1.8",
    "ISO": 200,
    "lighting": "Studio Lampe",
    "focal_point": "Beispiel Daten",
    "photo_genre": "Portrait"})}
user1 = user(images)
@app.route("/<id>")
def hello_world(id):
    global user1
    print(user1.images)
    print(id)
    best = {"location": user1.images[id].location, "info_there": user1.images[id].info_there, "info": user1.images[id].info}
    return jsonify(id)
@app.route("/add_image", methods=["POST"])
def save_image():
    global user1
    data = request.get_json() 
    # example data {'image_address': '1', 'Kameramodell': '2', 'Verwendetes_Objektiv': '3', 'Belichtungszeit': '4', 'Blende': '5', 'ISO': '6', 'Beleuchtung': '7', 'Fokuspunkt': '8', 'Fotogenre': '9'}
    image1 = image(str(uuid.uuid4()), data["image_address"], {}, {
    "camera_model": data["Kameramodell"],
    "lens_used": data["Verwendetes_Objektiv"],
    "shutter_speed": data["Belichtungszeit"],
    "aperture": data["Blende"],
    "ISO": data["ISO"],
    "lighting": data["Beleuchtung"],
    "focal_point": data["Fokuspunkt"],
    "photo_genre": data["Fotogenre"]})
    user1.images[image1.identificator] = image1
    print(user1.images)
    print(image1.identificator)
    return jsonify({"id": image1.identificator}), 200


if __name__ == "__main__":
    app.run(debug=True) 