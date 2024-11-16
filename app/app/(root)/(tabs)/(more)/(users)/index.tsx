import ListEmpty from "@/components/ListEmpty";
import { PagedResponse } from "@/models/pagedResponse";
import React, { useEffect, useState } from "react";
import { View } from "react-native";
import { SwipeListView } from "react-native-swipe-list-view";
import { ListItem, SwipeButton } from "@/components/ListItem";
import { router } from "expo-router";
import { fetchCustomers } from "@/api/customers";
import { CustomerResponse } from "@/models/customer/customerResponse.model";
import { fetchOrders } from "@/api/orders";
import { OrderResponse } from "@/models/order/orderResponse.model";
import { fetchUserList } from "@/api/user";
import { UserResponse } from "@/models/user/user.model";

const ItemsScreen = () => {
  const [data, setData] = useState<PagedResponse<UserResponse> | undefined>(undefined);

  useEffect(() => {
    getItems();
  }, []);

  const getItems = async () => {
    const response = await fetchUserList(1, 10);
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
              title={`${item.item.first_name} ${item.item.last_name} (${item.item.email})`}
              leftIconName="person-outline"
              onPress={() => console.log(item.item.id)}
            />
          )}
          renderHiddenItem={(item) => (
            <View style={{ flexDirection: "row", alignItems: "center", justifyContent: "flex-end", height: "100%" }}>
              <SwipeButton iconName="trash" onPress={() => console.log("Delete")} />
              <SwipeButton
                iconName="create"
                onPress={() => router.push({ pathname: "/(users)/user_form", params: { id: item.item.id } })}
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
