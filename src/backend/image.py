class image_base():
    def __init__(self, identificator: str, location: str, info_there: dict, info: dict):
        self.identificator = identificator  # always a uuid as a str
        self.location = location            # string
        self.info_there = info_there        # dict with bool as val
        self.info = info
    def debug(self):
        print("Identificator: " + self.identificator + "\nlocation: " + self.location + "\ninfo_there: " + str(self.info_there) + "\ninfo: " + str(self.info))
    def dictionary(self):
        image = {"identificator": self.identificator,
                 "location": self.location,
                 "info_there": self.info_there,
                 "info": self.info}
        return image
class user():
    def __init__(self, images: dict[str, image_base]):
        self.images = images # dict