import React from "react";
import {Avatar, AvatarImage, AvatarFallback} from "@/components/ui/avatar"

type UserItemProps = {
  imageSrc: string;
  alt: string;
  header: string;
  text: string;
};

const UserItem = ({imageSrc, alt, header, text}: UserItemProps) => {
  return (
    <div className="flex items-center justify-center gap-2 border rounded-[8px] p-2">
      <Avatar>
        <AvatarImage src={imageSrc} />
        <AvatarFallback>{alt}</AvatarFallback>
      </Avatar>
      <div>
        <p>{header}</p>
        <p>{text}</p>
      </div>
    </div>
  );
};

export default UserItem;
