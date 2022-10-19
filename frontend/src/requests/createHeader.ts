import { Header } from "@/models/Header"

function createHeader(headerType: string): Header {
    let header: Header;
    let contentType: string;
    if (headerType === "file") {
        contentType = "multipart/form-data";
    } else if (headerType === "json") {
        contentType = "application/json";
    } else if (headerType === "data") {
        contentType = "multipart/form-data";
    } else {
        contentType = "application/x-www-form-urlencoded";
    }
    header = {
        "Content-Type": contentType,
    };

    return header
}

export { createHeader }