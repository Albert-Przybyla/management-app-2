import { Link } from "expo-router";
import React from "react";
import { Text, View } from "react-native";

const home = () => {
  return (
    <View>
      <Text>Home</Text>
      <Link href="/(root)/(tabs)/coasd" className="mt-5">
        <Text>Go to home screen!</Text>
      </Link>
    </View>
  );
};

export default home;
