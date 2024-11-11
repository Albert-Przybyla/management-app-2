import { Stack } from "expo-router";
import React from "react";
import "@placekit/autocomplete-js/dist/placekit-autocomplete.css";

const Layout = () => {
  return (
    <Stack>
      <Stack.Screen name="(tabs)" options={{ headerShown: false }} />
    </Stack>
  );
};

export default Layout;
