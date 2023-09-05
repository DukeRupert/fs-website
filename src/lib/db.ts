import { createClient } from "@supabase/supabase-js";
import {
  PUBLIC_SUPABASE_URL,
  PUBLIC_SUPABASE_ANON_KEY,
} from "$env/static/public";
// import type { Database } from '$lib/types/supabase';

console.log(PUBLIC_SUPABASE_ANON_KEY, PUBLIC_SUPABASE_URL)

export const Supabase = createClient(
  PUBLIC_SUPABASE_URL,
  PUBLIC_SUPABASE_ANON_KEY
);
