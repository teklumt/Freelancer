import React from "react";
import { Button } from "@/components/ui/button";
import { Bell, Plus } from "lucide-react";
import Link from "next/link";
import { SignedIn, SignedOut, SignInButton } from "@clerk/nextjs";
import Image from "next/image";
const Navbar = () => {
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
              <Button variant={"secondary"}>
                Sign In
              </Button>
            </SignInButton>
          </SignedOut>
        </div>
      </div>
    </div>
  );
};

export default Navbar;
