import { Tabs } from "expo-router";
import React from "react";

const Layout = () => {
  return (
    <Tabs initialRouteName="index" screenOptions={{ tabBarActiveTintColor: "white" }}>
      <Tabs.Screen name="home" options={{ title: "Start", tabBarIcon: () => null }} />
      <Tabs.Screen name="cos" options={{ title: "Test", tabBarIcon: () => null }} />
      <Tabs.Screen name="items" options={{ title: "Zasoby", tabBarIcon: () => null }} />
    </Tabs>
  );
};

export default Layout;
