import axios from "axios";

export const mapBox = axios.create({
  baseURL: `${process.env.EXPO_PUBLIC_MAPBOX_URL}`,
  headers: {
    "Content-Type": "application/json",
  },
});
