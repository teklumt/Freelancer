"use client";
import { useRouter } from "next/navigation";
import React, { useEffect } from "react";
import { RedirectToSignIn, useAuth } from "@clerk/nextjs";

const Page = () => {
  const { isSignedIn } = useAuth();
  const router = useRouter();

  if (isSignedIn) {
    router.push("/client/protected");
    return null;
  } else {
    return <RedirectToSignIn />;
  }

  return <div>Client Page</div>;
};

export default Page;
