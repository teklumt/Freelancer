"use client"
import React, { useEffect } from 'react';
import { useAuth } from '@clerk/nextjs';
import { useRouter } from 'next/navigation';

const Page = () => {
  const { isSignedIn } = useAuth();
  const router = useRouter();

  useEffect(() => {
    if (isSignedIn) {
      router.push('/freelancer/protected');
    } else {
      router.push('/freelancer/unprotected');
    }
  }, [isSignedIn, router]);

  return (
    <div>
      {/* Optional: Loading state or empty content while redirecting */}
    </div>
  );
};

export default Page;
