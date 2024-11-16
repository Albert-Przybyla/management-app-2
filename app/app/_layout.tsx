import { useFonts } from "expo-font";
import { Stack } from "expo-router";
import * as SplashScreen from "expo-splash-screen";
import { useEffect } from "react";
import "react-native-reanimated";
import "../global.css";
import { AuthProvider } from "@/context/AuthContext";
import { GestureHandlerRootView } from "react-native-gesture-handler";
import { LocaleConfig } from "react-native-calendars";

SplashScreen.preventAutoHideAsync();

export default function RootLayout() {
  LocaleConfig.locales["pl"] = {
    monthNames: [
      "Styczeń",
      "Luty",
      "Marzec",
      "Kwiecień",
      "Maj",
      "Czerwiec",
      "Lipiec",
      "Sierpiń",
      "Wrzesien",
      "Październik",
      "Listopad",
      "Grudzień",
    ],
    monthNamesShort: ["Sty", "Lut", "Mar", "Kwi", "Maj", "Cze", "Lip", "Sie", "Wrz", "Paź", "Lis", "Gru"],
    dayNames: ["Poniedziałek", "Wtorek", "Sąbado", "Czwartek", "Piątek", "Sobota", "Niedziela"],
    dayNamesShort: ["Pon", "Wto", "Sob", "Czw", "Pią", "Sob", "Nie"],
    today: "Dzisiaj",
  };

  LocaleConfig.defaultLocale = "pl";

  const [loaded] = useFonts({
    "Jakarta-Bold": require("../assets/fonts/PlusJakartaSans-Bold.ttf"),
    "Jakarta-ExtraBold": require("../assets/fonts/PlusJakartaSans-ExtraBold.ttf"),
    "Jakarta-ExtraLight": require("../assets/fonts/PlusJakartaSans-ExtraLight.ttf"),
    "Jakarta-Light": require("../assets/fonts/PlusJakartaSans-Light.ttf"),
    "Jakarta-Medium": require("../assets/fonts/PlusJakartaSans-Medium.ttf"),
    "Jakarta-Regular": require("../assets/fonts/PlusJakartaSans-Regular.ttf"),
    "Jakarta-SemiBold": require("../assets/fonts/PlusJakartaSans-SemiBold.ttf"),
  });

  useEffect(() => {
    if (loaded) {
      SplashScreen.hideAsync();
    }
  }, [loaded]);

  if (!loaded) {
    return null;
  }

  return (
    <AuthProvider>
      <GestureHandlerRootView>
        <Layout />
      </GestureHandlerRootView>
    </AuthProvider>
  );
}

const Layout = () => {
  return (
    <Stack>
      <Stack.Screen name="+not-found" />
      {/* {authState?.authenticated ? (
        <Stack.Screen name="cos" options={{ headerShown: false }} />
      ) : (
        <Stack.Screen name="sign-in" options={{ headerShown: false }} />
      )} */}
      <Stack.Screen name="index" options={{ headerShown: false }} />
      <Stack.Screen name="(root)" options={{ headerShown: false }} />
      <Stack.Screen name="(auth)" options={{ headerShown: false }} />
    </Stack>
  );
};
