import { Avatar } from "@/components/ui/avatar";
import React from "react";
import UserItem from "./UserItem";
import { Button } from "@/components/ui/button";
import { Bell, Plus } from "lucide-react";
import Link from "next/link"
import { SignOutButton } from "@clerk/nextjs";
const Navbar = () => {
  return (
    <div className="flex gap-4 p-4 border-b justify-between">
      <div>
        <UserItem
          imageSrc="https://github.com/humaans/next-img/blob/master/assets/next-img.png"
          alt="Av"
          header="Company"
          text="Built to serve"
        />
      </div>

      <div className="flex gap-4 items-center">
        <div className="cursor-pointer">
          <Bell />
        </div>
        <div>
          <Button variant={"secondary"}>
            <Link href="/client/protected/post-job" className="flex gap-2">
              <Plus className="mr-2 h-4 w-4" /> Post a Job
            </Link>
          </Button>
        </div>
        <div>
          <SignOutButton>
            <Button variant={"secondary"}>
              Sign Out
            </Button>
          </SignOutButton>
        </div>
      </div>
    </div>
  );
};

export default Navbar;
