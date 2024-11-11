import { authConstants } from "@/constants/auth";
import { auth } from "@/firebaseConfig";
import { secureGetValueFor, secureSave } from "@/utils/store";
import {
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
} from "firebase/auth";

export async function signup(email: string, password: string) {
  try {
    const user = await createUserWithEmailAndPassword(auth, email, password);
    return user.user.getIdToken();
  } catch (error) {
    return error;
  }
}

export async function login(email: string, password: string) {
  try {
    const user = await signInWithEmailAndPassword(auth, email, password);
    return user.user.getIdToken();
  } catch (error) {
    return error;
  }
}

export async function logout() {
  try {
    await auth.signOut();
  } catch (error) {
    console.log(error);
  }
}
export async function storeToken(token: string) {
  try {
    await secureSave(authConstants.AUTH_TOKEN_KEY, token);
  } catch (error) {
    console.log(error);
  }
}

export async function getAuthToken() {
  try {
    const token = await secureGetValueFor(authConstants.AUTH_TOKEN_KEY);
    return token;
  } catch (error) {
    console.log(error);
  }
}
