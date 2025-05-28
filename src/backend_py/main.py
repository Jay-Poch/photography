import uuid
from flask import Flask, jsonify, request
from flask_cors import CORS # type: ignore
from image import image_base, user
app = Flask(__name__)
_ = CORS(app)
images = {"example": image_base("example", "	http://0.0.0.0:2020/images/866-536x354.jpg", {"fstop": True}, {
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
    return jsonify(user1.images[id].dictionary())


@app.route("/add_image", methods=["POST"])
def save_image():
    global user1
    data = request.get_json() 
    image1 = image_base(str(uuid.uuid4()), data["image_address"], {}, {
    "camera_model": data["Kameramodell"],
    "lens_used": data["Verwendetes_Objektiv"],
    "shutter_speed": data["Belichtungszeit"],
    "aperture": data["Blende"],
    "ISO": data["ISO"],
    "lighting": data["Beleuchtung"],
    "focal_point": data["Fokuspunkt"],
    "photo_genre": data["Fotogenre"]})
    user1.images[image1.identificator] = image1
    return jsonify({"id": image1.identificator}), 200


if __name__ == "__main__":
    app.run(debug=True) 