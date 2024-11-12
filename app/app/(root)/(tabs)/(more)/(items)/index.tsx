import { fetchItems } from "@/api/items";
import ListEmpty from "@/components/ListEmpty";
import ListHeader from "@/components/ListHeader";
import ParallaxScrollView from "@/components/ParallaxScrollView";
import { ItemResponse } from "@/models/item/itemResponse.model";
import { PagedResponse } from "@/models/pagedResponse";
import React, { useEffect, useLayoutEffect, useState } from "react";
import { Button, FlatList, Image, Text, View } from "react-native";
import Animated, { useAnimatedRef } from "react-native-reanimated";
import { useNavigation } from "@react-navigation/native";
import { SwipeListView } from "react-native-swipe-list-view";
import { ListItem, SwipeButton } from "@/components/ListItem";
import { router } from "expo-router";

const ItemsScreen = () => {
  const [data, setData] = useState<PagedResponse<ItemResponse> | undefined>(undefined);
  const navigation = useNavigation();

  useEffect(() => {
    getItems();
  }, []);

  const getItems = async () => {
    const response = await fetchItems(1, 10);
    if (response) {
      setData(response);
    }
  };

  return (
    <View>
      {data && (
        <SwipeListView
          className="h-full"
          keyExtractor={(item) => item.id}
          data={data.items}
          renderItem={(item) => (
            <ListItem
              title={item.item.name}
              leftIconName="construct-outline"
              onPress={() => console.log(item.item.id)}
            />
          )}
          renderHiddenItem={(item) => (
            <View style={{ flexDirection: "row", alignItems: "center", justifyContent: "flex-end", height: "100%" }}>
              <SwipeButton iconName="trash" onPress={() => console.log("Delete")} />
              <SwipeButton
                iconName="create"
                onPress={() => router.push({ pathname: "/(items)/item_form", params: { id: item.item.id } })}
              />
            </View>
          )}
          ListEmptyComponent={ListEmpty}
          rightOpenValue={-130}
          disableRightSwipe
        />
      )}
    </View>
  );
};

export default ItemsScreen;
