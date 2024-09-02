"use client";
import { useEffect } from "react";
import { useRouter } from "next/navigation";

import {
  SignedIn,
  SignedOut,
  SignInButton,
  SignOutButton,
  useAuth,
} from "@clerk/nextjs";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import Image from "next/image";
const Navbar = () => {
  const { isSignedIn } = useAuth();
  const router = useRouter();

  useEffect(() => {
    if (isSignedIn) {
      router.push("/freelancer/protected");
    }
  }, [isSignedIn, router]);
  return (
    <div className="flex gap-4 p-4 border-b justify-between">
      <div className="flex gap-10">
        <div>
          <Image src={"/logo.png"} alt={"Logo"} width={160} height={36}></Image>
        </div>
        <div className="flex gap-4 items-center">
          <div>
            <Link href={"/freelancer/unprotected/find-jobs"}>
              <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
                Find Jobs
              </h4>
            </Link>
          </div>
          <div>
            <Link href={"/freelancer/unprotected/browse-companies"}>
              <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
                Browse Companies
              </h4>
            </Link>
          </div>
        </div>
      </div>
      <div className="flex gap-4 items-center">
        <div>
          <SignedOut>
            <SignInButton>
              <Button variant={"secondary"}>Sign In</Button>
            </SignInButton>
          </SignedOut>
          <SignedIn>
            <SignOutButton>
              <Button variant={"secondary"}>Sign Out</Button>
            </SignOutButton>
          </SignedIn>
        </div>
      </div>
    </div>
  );
};

export default Navbar;
