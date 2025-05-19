from flask import Flask, jsonify, request
import json
from flask_cors import CORS # type: ignore
from image import image, user
app = Flask(__name__)
_ = CORS(app)
example_image = {"example": image("example", "/media/jay/29e62c74-c999-43f4-aa58-27294a6f1dc9/11_my_programming_projects/real/photography/866-536x354.jpg", {"fstop": True}, {
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
if __name__ == "__main__":
    app.run(debug=True) 
    
@app.route("/add_image", methods=["POST"])
def save_image():
    data = request.get_json()
    
    # ✅ Example: Save to file (you can use a DB instead)
    with open("data.json", "a") as f:
        f.write(json.dumps(data) + "\n")

    return jsonify({"status": "success"}), 200
