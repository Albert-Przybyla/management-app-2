import { fetchItems } from "@/api/items";
import ListEmpty from "@/components/ListEmpty";
import ListHeader from "@/components/ListHeader";
import ParallaxScrollView from "@/components/ParallaxScrollView";
import { ItemResponse } from "@/models/item/itemResponse.model";
import { PagedResponse } from "@/models/pagedResponse";
import React, { useEffect, useState } from "react";
import { FlatList, Image, Text, View } from "react-native";
import Animated, { useAnimatedRef } from "react-native-reanimated";

const listHeader = () => {
  return (
    <View>
      <Text>Header</Text>
    </View>
  );
};

const ItemsScreen = () => {
  const [data, setData] = useState<PagedResponse<ItemResponse> | undefined>(undefined);
  const scrollRef = useAnimatedRef<Animated.FlatList<any>>();
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
    <View>
      <Text>{item.name}</Text>
      <Text>{item.description}</Text>
    </View>
  );

  return (
    <View>
      {data?.items && (
        <Animated.FlatList
          ref={scrollRef}
          className="h-full"
          data={data.items}
          renderItem={renderItem}
          keyExtractor={(item) => item.id}
          ListHeaderComponent={<ListHeader scrollRef={scrollRef} headerImage={<Text>Header Image</Text>} />}
          ListEmptyComponent={ListEmpty}
        />
      )}
    </View>
  );
};

export default ItemsScreen;
