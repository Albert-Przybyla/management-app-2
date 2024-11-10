import { Text as RNText, type TextProps as RNTextProps, StyleSheet } from "react-native";

export type TextProps = RNTextProps & {
  type?: "default" | "title" | "defaultSemiBold" | "subtitle" | "link";
};

export function Text({ style, type = "default", ...rest }: TextProps) {
  return (
    <RNText
      style={[
        type === "default" ? styles.default : undefined,
        type === "title" ? styles.title : undefined,
        type === "defaultSemiBold" ? styles.defaultSemiBold : undefined,
        type === "subtitle" ? styles.subtitle : undefined,
        type === "link" ? styles.link : undefined,
        style,
      ]}
      className="font-jakartaRegular"
      {...rest}
    />
  );
}

const styles = StyleSheet.create({
  default: {
    fontSize: 16,
    lineHeight: 24,
  },
  defaultSemiBold: {
    fontSize: 16,
    lineHeight: 24,
    fontWeight: "600",
  },
  title: {
    fontSize: 18,
    fontWeight: "600",
    lineHeight: 32,
  },
  subtitle: {
    fontSize: 20,
    fontWeight: "600",
  },
  link: {
    lineHeight: 30,
    fontSize: 16,
    color: "#0a7ea4",
  },
});
