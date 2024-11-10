import ListEmpty from "@/components/ListEmpty";
import { ListItem, SwipeButton } from "@/components/ListItem";
import { Text, View } from "react-native";
import { SwipeListView } from "react-native-swipe-list-view";

const MoreScreen = () => {
  const moreScreens = [
    {
      name: "Items",
      path: "/(root)/(tabs)/(more)/items",
      icon: "search-outline",
    },
    {
      name: "Cos",
      path: "/(root)/(tabs)/(more)/cos",
      icon: "search-outline",
    },
  ];

  return (
    <View>
      <SwipeListView
        className="h-full"
        keyExtractor={(item) => item.name}
        data={moreScreens}
        renderItem={(item) => (
          <ListItem
            title={item.item.name}
            leftIconName="star"
            rightIconName="chevron-forward"
            onPress={() => console.log(item.item.name)}
          />
        )}
        renderHiddenItem={(item) => (
          <View style={{ flexDirection: "row", alignItems: "center", justifyContent: "flex-end", height: "100%" }}>
            <SwipeButton iconName="trash" onPress={() => console.log("Delete")} />
            <SwipeButton iconName="create" onPress={() => console.log("Edit")} />
          </View>
        )}
        ListEmptyComponent={ListEmpty}
        rightOpenValue={-130}
        disableRightSwipe
      />
    </View>
  );
};

export default MoreScreen;
