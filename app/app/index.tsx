import { Redirect } from "expo-router";
import React from "react";
import { SafeAreaView, Text } from "react-native";

const index = () => {
  return <Redirect href="/(auth)/welcome" />;
};

export default index;
