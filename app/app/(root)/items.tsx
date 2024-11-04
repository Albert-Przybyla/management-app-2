import { fetchItems } from "@/api/items";
import ParallaxTitleScrollView from "@/components/ParallaxTitleScrollView";
import { ThemedText } from "@/components/ThemedText";
import { ThemedView } from "@/components/ThemedView";
import { ItemResponse } from "@/models/item/itemResponse.model";
import { PagedResponse } from "@/models/pagedResponse";
import React, { useEffect, useState } from "react";
import { FlatList, ListRenderItem, View, VirtualizedList } from "react-native";

const ItemsScreen = () => {
  const [data, setData] = useState<PagedResponse<ItemResponse> | undefined>(undefined);

  useEffect(() => {
    getItems();
  }, []);

  const getItems = async () => {
    const response = await fetchItems(1, 10);
    if (response) {
      setData(response);
    }
  };

  const renderItem = ({ item }: { item: ItemResponse }) => (
    <ThemedView>
      <ThemedText type="subtitle">{item.name}</ThemedText>
      <ThemedText>{item.description}</ThemedText>
    </ThemedView>
  );

  return (
    <View>
      {data?.items && <FlatList data={data.items} renderItem={renderItem} keyExtractor={(item) => item.id} />}
    </View>
  );
};

export default ItemsScreen;
