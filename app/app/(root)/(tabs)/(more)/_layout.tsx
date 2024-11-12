import { router, Stack } from "expo-router";
import React from "react";
import { Button } from "react-native";

const Layout = () => {
  return (
    <Stack screenOptions={{ title: "Więcej" }}>
      <Stack.Screen name="index" />
      <Stack.Screen
        name="(items)"
        options={{
          headerShown: true,
          headerTitle: "Zasoby",
          headerBackTitle: "Cofnij",
          headerRight: () => (
            <Button
              onPress={() => {
                router.push({
                  pathname: "/(items)/item_form",
                  params: { id: undefined },
                });
              }}
              title="Dodaj"
            />
          ),
        }}
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
          headerRight: () => (
            <Button
              onPress={() => {
                console.log("Add item");
              }}
              title="Dodaj"
            />
          ),
          presentation: "modal",
        }}
        name="cos"
      />
    </Stack>
  );
};

export default Layout;
