import { Button } from "@/components/ui/button";
import Navbar from "../client/components/Navbar";
import Sidebar from "../client/components/Sidebar";
import { SignInButton, SignedIn, SignedOut, UserButton } from "@clerk/nextjs";
import Link from "next/link";
export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <html lang="en">
        <div className="flex justify-between items-start">
          <aside>
            <Sidebar></Sidebar>
          </aside>
          <main className="flex flex-col w-full h-full pl-[300px]">
            <header>
              <Navbar></Navbar>
            </header>
            <div className="">{children}</div>
          </main>
        </div>
      </html>
    </>
  );
}
