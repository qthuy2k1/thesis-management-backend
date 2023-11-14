import { IGeneralLinkAttachment } from "../interface/submit";

const { google } = require("googleapis");
const path = require("path");
const fs = require("fs");
const CLIENT_ID =
  "1026044901751-hfanduani1dbdbfleravpcf9vco9v0sm.apps.googleusercontent.com";
const CLIENT_SECRET = "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";
const REDIRECT_URI = "https://developers.google.com/oauthplayground";

const REFRESH_TOKEN =
  "1//04Smt9XvC-QlFCgYIARAAGAQSNwF-L9IrbR7PJMjvlDVQqxnTPsNqlZmeHRaIxWxyxr6UiEXl_hrrYqLEGyoSIMbVTWB5eyIonyw";

const oauth2Client = new google.auth.OAuth2(
  CLIENT_ID,
  CLIENT_SECRET,
  REDIRECT_URI
);

oauth2Client.setCredentials({ refresh_token: REFRESH_TOKEN });

const drive = google.drive({
  version: "v3",
  auth: oauth2Client,
});

export async function uploadFile(file: File, pathFile: string) {
  const response = await drive.files.create({
    requestBody: {
      name: file,
      mimeType: file.type,
    },
    media: {
      mimeType: file.type,
      body: fs.createReadStream(pathFile),
    },
  });
  return response.data.id;
}

export async function generatePublicUrl(
  id: string
): Promise<IGeneralLinkAttachment> {
  try {
    const fileId = id;
    await drive.permissions.create({
      fileId: fileId,
      requestBody: {
        role: "reader",
        type: "anyone",
      },
    });

    const result = await drive.files.get({
      fileId: fileId,
      fields: "webViewLink, name",
    });

    const linkAttachment: IGeneralLinkAttachment = {
      id: fileId,
      name: result.data.name,
      src: result.data.webViewLink,
    };

    return linkAttachment;
  } catch (error) {
    console.log(error);
    throw error;
  }
}

export async function uploadAndGeneratePublicUrl(file: File, pathFile: string) {
  const initiate = { id: "", name: "", src: "" };
  try {
    const fileId = await uploadFile(file, pathFile);
    const publicUrl = await generatePublicUrl(fileId);
    return publicUrl;
  } catch (error) {
    console.log(error);
    return initiate;
  }
}

// export async function deleteFile() {
//   try {
//     const response = await drive.files.delete({
//       fileId: "YOUR FILE ID",
//     });
//     console.log(response.data, response.status);
//   } catch (error) {
//     console.log(error);
//   }
// }
