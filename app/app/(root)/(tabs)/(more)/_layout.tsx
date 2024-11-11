import { TabBarIcon } from "@/components/TabBarIcon";
import { Stack, Tabs } from "expo-router";
import React from "react";

const Layout = () => {
  return (
    <Stack screenOptions={{ title: "Więcej" }}>
      <Stack.Screen name="index" />
      <Stack.Screen
        options={{
          headerShown: true,
          headerTitle: "Zasoby",
          headerBackTitle: "Cofnij",
        }}
        name="items"
      />
      <Stack.Screen
        options={{
          headerShown: true,
          headerTitle: "Klienci",
          headerBackTitle: "Cofnij",
        }}
        name="customers"
      />
      <Stack.Screen
        options={{
          headerShown: true,
          headerTitle: "Zasoby",
          headerBackTitle: "Cofnij",
        }}
        name="cos"
      />
    </Stack>
  );
};

export default Layout;
