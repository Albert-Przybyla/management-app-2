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

export const isTokenValid = async (token: string): Promise<boolean> => {
  const decodedToken = await getDecodedToken(token);
  if (decodedToken && decodedToken.exp) {
    const currentTime = Math.floor(Date.now() / 1000);
    return decodedToken.exp > currentTime;
  }
  return false;
};
