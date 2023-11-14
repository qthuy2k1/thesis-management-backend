import multer from "multer";
import fs from "fs";

const storage = multer.diskStorage({
  destination: function (req, file, cb) {
    const post = req.body;
    const uploadPath = `src/uploads/${post.uid}`;
    fs.access(uploadPath, (error) => {
      if (error) {
        fs.mkdir(uploadPath, { recursive: true }, (error) => {
          if (error) {
            console.error("Error creating upload directory:", error);
            cb(error, "err");
          } else {
            cb(null, uploadPath);
          }
        });
      } else {
        cb(null, uploadPath);
      }
    });
  },
  filename: function (req, file, cb) {
    cb(null, file.originalname);
  },
});

const uploadMiddleware = multer({ storage: storage });

export default uploadMiddleware;