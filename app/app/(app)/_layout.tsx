import { Stack } from "expo-router";

export const unstable_settings = {
  initialRouteName: "(root)",
};

export default function AppLayout() {
  return (
    <Stack>
      <Stack.Screen
        name="login"
        options={{
          presentation: "modal",
        }}
      />
      <Stack.Screen name="(tabs)" />
    </Stack>
  );
}
