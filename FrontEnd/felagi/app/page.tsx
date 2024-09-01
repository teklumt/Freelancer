"use client";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { SignInButton, SignedOut } from "@clerk/nextjs";

export default function Home() {
  return (
    <main>
      <SignedOut>
        <SignInButton>
          <Button variant={"outline"} asChild>
            <Link href="./client">Client</Link>
          </Button>
        </SignInButton>
      </SignedOut>
      <Button variant={"outline"} asChild>
        <Link href="./freelancer">Freelancer</Link>
      </Button>
    </main>
  );
}
