"use client";
import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function Home() {
  return (
    <main>
      <Button variant={"outline"}>
        <Link href="./client">Client</Link>
      </Button>
      <Button variant={"outline"} asChild>
        <Link href="./freelancer">Freelancer</Link>
      </Button>
    </main>
  );
}
