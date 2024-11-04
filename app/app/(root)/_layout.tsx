import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { createDrawerNavigator, DrawerContentScrollView, DrawerItem } from "@react-navigation/drawer";
import { NavigationContainer } from "@react-navigation/native";
import HomeScreen from ".";
import ItemsScreen from "./items";
import TabTwoScreen from "./explore";

const Drawer = createDrawerNavigator();

export default function TabLayout() {
  return (
    <>
      <Drawer.Navigator drawerContent={(props) => <CustomDrawerContent {...props} />}>
        <Drawer.Screen name="Home" component={HomeScreen} options={{ title: "Home" }} />
        <Drawer.Screen name="Items" component={ItemsScreen} options={{ title: "Zasoby" }} />
        <Drawer.Screen name="Profile" component={TabTwoScreen} options={{ title: "Profile" }} />
      </Drawer.Navigator>
    </>
  );
}

function CustomDrawerContent(props: any) {
  // Przykładowe dane użytkownika
  const userName = "Albert Przybyła";
  const userEmail = "albert@example.com";

  const handleLogout = () => {
    // Logika wylogowania
    console.log("User logged out");
    // Tutaj możesz dodać funkcję wylogowania lub nawigować do ekranu logowania
  };

  return (
    <DrawerContentScrollView {...props}>
      {/* Standardowe elementy drawer */}
      <DrawerItem label="Home" onPress={() => props.navigation.navigate("Home")} />
      <DrawerItem label="Zasoby" onPress={() => props.navigation.navigate("Items")} />
      <DrawerItem label="Profile" onPress={() => props.navigation.navigate("Profile")} />

      {/* Dolna sekcja z danymi użytkownika i przyciskiem wylogowania */}
      <View style={styles.footer}>
        <Text style={styles.userInfo}>{userName}</Text>
        <Text style={styles.userInfo}>{userEmail}</Text>
        <DrawerItem label="Wyloguj" onPress={handleLogout} style={styles.logoutButton} />
      </View>
    </DrawerContentScrollView>
  );
}

const styles = StyleSheet.create({
  footer: {
    padding: 20,
    borderTopWidth: 1,
    borderColor: "#ddd",
  },
  userInfo: {
    fontSize: 16,
    fontWeight: "bold",
    marginVertical: 4,
  },
  logoutButton: {
    marginTop: 10,
    backgroundColor: "#ff4d4f",
  },
});
