import xml2js from "xml2js";

// Function to transform a xml string to json object
export function parseXML(data) {
  return new Promise((resolve) => {
    let k = "";
    let arr = [],
      parser = new xml2js.Parser({
        trim: true,
        explicitArray: true,
      });
    parser.parseString(data, function (err, result) {
      console.log("result", result);
      let obj = result.rss.channel[0].item;
      console.log(obj);
      for (k in obj) {
        console.log(obj[k].title[0]);
        console.log(obj[k].description[0]);
      }
      resolve(arr);
    });
  });
}
