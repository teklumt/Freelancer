"use client"
import UserItem from "./UserItem";
import { Button } from "@/components/ui/button";
import { Bell, Plus } from "lucide-react";
import Link from "next/link";
import { SignedIn, SignOutButton, useAuth } from "@clerk/nextjs";
import { useRouter } from "next/navigation";
import { useEffect } from "react";
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
          <SignOutButton>
            <Button variant={"secondary"}>
              <Link
                href="/freelancer/unprotected/landing-page"
                className="flex gap-2"
              >
                <Plus className="mr-2 h-4 w-4" /> Back to Homepage
              </Link>
            </Button>
          </SignOutButton>
        </div>
        <SignedIn>
          <SignOutButton>
            <Button variant={"secondary"}>Sign Out</Button>
          </SignOutButton>
        </SignedIn>
      </div>
    </div>
  );
};

export default Navbar;
