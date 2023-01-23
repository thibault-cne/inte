import { Stars } from "./Stars";

type User = {
  id: string;
  last_name: string;
  first_name: string;
  email: string;
  current_year: number;
  promotion_year: number;
  points: number;
  god_father_id: number;
  facebook_id: string;
  google_id: string;
  instagram_id: string;
  snapchat_id: string;
  hometown: string;
  studies: string;
  personal_description: string;
  profile_picture: string;
  last_login: string;
  user_type: string;
  color: string;
  stars: Stars[];
};

function getName(user: User): string {
  return user.first_name + " " + user.last_name;
}

export { User, getName };
