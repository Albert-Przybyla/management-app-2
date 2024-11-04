import { DecodedToken } from "@/models/user/decodedToken.model";
import { User } from "@/models/user/user.model";
import { getDecodedToken } from "@/services/authService";
import { useRouter } from "expo-router";
import * as SecureStore from "expo-secure-store";
import { createContext, ReactNode, useContext, useEffect, useState } from "react";

interface UserContextType {
  user: User | null;
  decodedToken: DecodedToken | null;
  login: (token: string) => Promise<void>;
  logout: () => void;
  isAuthenticated: boolean;
  loading: boolean;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

interface UserProviderProps {
  children: ReactNode;
}

export const UserProvider = ({ children }: UserProviderProps) => {
  const [user, setUser] = useState<User | null>(null);
  const [token, setToken] = useState<string | null>(null);
  const [decodedToken, setDecodedToken] = useState<DecodedToken | null>(null);

  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const initializeToken = async () => {
      const storedToken = await SecureStore.getItemAsync("token");
      if (storedToken) {
        setToken(storedToken);
      } else {
        setToken(null);
      }
      setLoading(false);
    };
    initializeToken();
  }, []);

  useEffect(() => {
    if (token) {
      decodeToken(token);
    }
  }, [token]);

  const decodeToken = async (token: string) => {
    const tokenData = await getDecodedToken(token);
    if (tokenData) {
      setDecodedToken(tokenData);
    } else {
    }
  };

  const login = async (token: string) => {
    try {
      //   await AsyncStorage.setItem("authToken", token);
      //   const userData = await getUserData(token);
      //   setUser(userData);
    } catch (error) {
      console.error("Błąd podczas logowania:", error);
    }
  };

  const logout = async () => {
    setUser(null);
  };

  const isAuthenticated = user !== null;

  return (
    <UserContext.Provider value={{ user, decodedToken, login, logout, isAuthenticated, loading }}>
      {children}
    </UserContext.Provider>
  );
};

export const useUser = () => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error("useUser must be used within a UserProvider");
  }
  return context;
};
