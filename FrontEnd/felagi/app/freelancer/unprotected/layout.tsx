import { SignedIn, SignedOut } from "@clerk/nextjs";
import Navbar from "./components/Navbar";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="p-4">
        <SignedOut>
          <header className="w-full">
            <Navbar></Navbar>
          </header>
          <main>{children}</main>
        </SignedOut>
        <SignedIn>
          <main>{children}</main>
        </SignedIn>
      </body>
    </html>
  );
}
