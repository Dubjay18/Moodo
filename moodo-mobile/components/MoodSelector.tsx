import React from "react";
import { Pressable, View } from "react-native";
import { ThemedText } from "./ThemedText";
import { ThemedView } from "./ThemedView";
import { useThemeColor } from "@/hooks/useThemeColor";
import { Colors } from "@/constants/Colors";
import { useAnimatedStyle, useSharedValue } from "react-native-reanimated";

const SIZE = 100;
function MoodSelector() {
  const translateX = useSharedValue(0);
  const translateY = useSharedValue(0);
  const contextX = useSharedValue(0);
  const contextY = useSharedValue(0);
  const progress = useSharedValue(1);
  const scale = useSharedValue(2);
  const color = useSharedValue("blue");

  // const {} = useThemeColor()
  const reanimatedStyle = useAnimatedStyle(() => {
    return {
      borderRadius: (progress.value * SIZE) / 2,
      backgroundColor: color.value,
      transform: [
        { translateX: translateX.value },
        { translateY: translateY.value },
        { scale: scale.value },
        {
          rotate: `${progress.value * 2 * Math.PI}rad`,
        },
      ],
    };
  });
  return (
    <Pressable
      onPress={() => {
        console.log("sup");
      }}
    >
      <ThemedView
        darkColor={Colors.dark.foreground}
        className="bg-red-300 h-20 w-20"
        style={{
          height: 100,
          width: 100,
          //   backgroundColor: "red",
          borderRadius: 50,
          alignItems: "center",
          justifyContent: "center",
          flexDirection: "row",
          display: "flex",
        }}
      >
        <ThemedText
          style={{
            marginEnd: "auto",
            marginStart: "auto",
          }}
        >
          mood
        </ThemedText>
      </ThemedView>
    </Pressable>
  );
}

export default MoodSelector;
