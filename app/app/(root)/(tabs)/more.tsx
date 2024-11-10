import ListEmpty from "@/components/ListEmpty";
import { ListItem } from "@/components/ListItem";
import { Text } from "@/components/Text";
import { useAuth } from "@/context/AuthContext";
import { Ionicons } from "@expo/vector-icons";
import { router } from "expo-router";
import { SafeAreaView, SectionList, View } from "react-native";

const MoreScreen = () => {
  const { authState, onLogout } = useAuth();
  const moreScreens = [
    {
      title: "Organizacja",
      data: [
        {
          name: "Items",
          icon: "search-outline",
          action: () => router.push("/(root)/(tabs)/(more)/items"),
        },
        {
          name: "Cos",
          icon: "search-outline",
          action: () => router.push("/(root)/(tabs)/(more)/cos"),
        },
      ],
    },
    {
      title: "Użytkownik",
      data: [
        {
          name: "Items",
          icon: "search-outline",
          action: () => router.push("/(root)/(tabs)/(more)/items"),
        },
        {
          name: "Wyloguj sie",
          icon: "search-outline",
          action: () => onLogout!(),
        },
      ],
    },
  ];

  const renderHeader = () => {
    return (
      <View className="p-2 space-x-2 w-full bg-white flex flex-row items-center border-b border-neutral-200">
        <Ionicons name="person" size={42} style={{ marginRight: 8 }} />
        <View className="p-2">
          <Text type="title">
            {authState?.userData?.first_name} {authState?.userData?.last_name}
          </Text>
          <Text type="subtitle">{authState?.userData?.email}</Text>
        </View>
      </View>
    );
  };

  return (
    <SafeAreaView>
      <SectionList
        className="h-full"
        keyExtractor={(item) => item.name}
        sections={moreScreens}
        renderItem={(item) => (
          <ListItem
            title={item.item.name}
            leftIconName="star"
            rightIconName="chevron-forward"
            onPress={() => () => item.item.action()}
          />
        )}
        ListEmptyComponent={ListEmpty}
        ListHeaderComponent={renderHeader}
        renderSectionHeader={({ section }) => <Text className="p-2 bg-neutral-100">{section.title}</Text>}
      />
    </SafeAreaView>
  );
};

export default MoreScreen;
