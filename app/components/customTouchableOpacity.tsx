import { ButtonProps } from "@/types/type";
import React from "react";
import { Text, TouchableOpacity } from "react-native";

const CoustomTouchableOpacity = ({ title, onPress, LeftIcon, RightIcon, className, ...props }: ButtonProps) => {
  return (
    <TouchableOpacity
      onPress={onPress}
      className={`w-full rounded-full flex flex-row items-center justify-center shadow-md shadow-neutral-400/70 ${className}`}
      {...props}
    >
      {LeftIcon && <LeftIcon />}
      <Text>{title}</Text>
      {RightIcon && <RightIcon />}
    </TouchableOpacity>
  );
};

export default CoustomTouchableOpacity;
