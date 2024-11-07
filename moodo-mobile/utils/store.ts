import * as SecureStore from "expo-secure-store";

export async function secureSave(key: string, value: string) {
  await SecureStore.setItemAsync(key, value);
}

export async function secureGetValueFor(key: string) {
  let result = await SecureStore.getItemAsync(key);
  return result;
}