import { DecodedToken } from "@/models/user/decodedToken.model";
import * as SecureStore from "expo-secure-store";
import { jwtDecode } from "jwt-decode";

export const getDecodedToken = async (token: string): Promise<DecodedToken | null> => {
  if (token) {
    try {
      return jwtDecode<DecodedToken>(token);
    } catch (error) {
      console.error("Invalid token:", error);
      return null;
    }
  }
  return null;
};
