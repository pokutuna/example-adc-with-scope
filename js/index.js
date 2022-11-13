const { google } = require("googleapis");

const spreadsheetId = process.env["SPREADSHEET_ID"];
const range = process.env["SHEET_RANGE"];

exports.app = async (_req, res) => {
  const auth = new google.auth.GoogleAuth({
    scopes: "https://www.googleapis.com/auth/spreadsheets",
  });

  const sheets = google.sheets({ version: "v4", auth });

  // read current
  const result = await sheets.spreadsheets.values.get({ spreadsheetId, range });

  const count = parseInt(result.data.values[0][0], 10);

  // update
  const updated = await sheets.spreadsheets.values.update({
    spreadsheetId,
    range,
    requestBody: {
      values: [[count + 1, new Date().toISOString(), "js"]],
    },
    valueInputOption: "USER_ENTERED",
    includeValuesInResponse: true,
  });

  return res.json({
    from: result.data.values[0],
    to: updated.data.updatedData.values[0],
  });
};
