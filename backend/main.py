from flask import Flask, jsonify, request
import json
from flask_cors import CORS # type: ignore
from image import image, user
app = Flask(__name__)
_ = CORS(app)
example_image = {"example": image("example", "	http://0.0.0.0:2020/images/866-536x354.jpg", {"fstop": True}, {
    "camera_model": "Canon EOS R5",
    "lens_used": "RF 50mm f/1.2L USM",
    "shutter_speed": "1/250 sec",
    "aperture": "f/1.8",
    "ISO": 200,
    "lighting": "Studio Lampe",
    "focal_point": "Beispiel Daten",
    "photo_genre": "Portrait"})}
user = user(example_image)
@app.route("/<id>")
def hello_world(id):
    best = {"location": user.images[id].location, "info_there": user.images[id].info_there, "info": user.images[id].info}
    return jsonify(best)
@app.route("/add_image", methods=["POST"])
def save_image():
    global user
    data = request.get_json()
    print(data)
    

    return jsonify({"status": "success"}), 200


if __name__ == "__main__":
    app.run(debug=True) 