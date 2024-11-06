import { SafeAreaView, StatusBar, Text } from "react-native";

export default function HomeScreen() {
  return (
    <SafeAreaView className={"flex-1 items-center justify-start"}>
      <Text className="text-red-900">Hello</Text>
      <StatusBar style="auto" />
    </SafeAreaView>
  );
}
