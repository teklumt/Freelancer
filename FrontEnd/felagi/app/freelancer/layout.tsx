"use client";
import { useRouter } from "next/navigation";
import { useEffect } from "react";
import Navbar from "./components/Navbar";
import Sidebar from "./components/Sidebar";
import { SignedIn, SignedOut } from "@clerk/nextjs";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();

  useEffect(() => {
    // Redirect only on the client side when SignedOut is detected
    if (typeof window !== "undefined") {
      router.push("/freelancer/unprotected");
    }
  }, [router]);

  return (
    <>
      <SignedIn>
        <html lang="en">
          <body>
            <div className="flex justify-between items-start">
              <aside>
                <Sidebar />
              </aside>
              <main className="flex flex-col w-full h-full pl-[300px]">
                <header>
                  <Navbar />
                </header>
                <div>{children}</div>
              </main>
            </div>
          </body>
        </html>
      </SignedIn>
      <SignedOut>
        <div>{children}</div>
      </SignedOut>
    </>
  );
}
