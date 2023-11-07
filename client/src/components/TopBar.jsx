import { Button } from "@mui/material";
import { Link } from "react-router-dom";

const TopBar = () => {
  return (
    <div className="flex space-x-5 p-5">
      <svg
        width="200"
        height="40"
        viewBox="0 0 160 20"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          fill-rule="evenodd"
          clip-rule="evenodd"
          d="M39.8294 10.0034C39.8331 10.6246 39.7285 11.2418 39.5204 11.8281C39.326 12.3261 39.0333 12.7811 38.6592 13.1667C38.2886 13.5401 37.8393 13.8288 37.3432 14.0123C36.3505 14.3786 35.26 14.3956 34.2561 14.0603C33.8126 13.9035 33.4045 13.6627 33.0547 13.3516C32.7017 13.0204 32.4082 12.6326 32.1866 12.2047C31.9497 11.7161 31.8027 11.1898 31.7525 10.6505H30.5371V14.139H28.6273V5.86443H30.5198V9.14755H31.7872C31.8498 8.63144 32.0025 8.12991 32.2387 7.66518C32.4571 7.25642 32.753 6.89268 33.1102 6.59363C33.4648 6.28758 33.8781 6.05476 34.3256 5.90894C34.792 5.75879 35.2795 5.68252 35.7701 5.68299C36.3172 5.68242 36.8602 5.77627 37.3744 5.96029C37.8642 6.14188 38.3084 6.42566 38.6766 6.79219C39.0502 7.18208 39.3428 7.64032 39.5377 8.14105C39.7581 8.73136 39.863 9.35749 39.8468 9.98631L39.8294 10.0034ZM19.9809 9.07566C19.9677 9.78408 19.8899 10.4899 19.7482 11.1845C19.6715 11.536 19.5134 11.8652 19.2864 12.1465C19.2008 12.2504 19.0913 12.3326 18.9669 12.3862C18.8438 12.4371 18.7115 12.4627 18.578 12.4615H18.3245V14.1732H18.8385C19.3746 14.1962 19.9019 14.0326 20.3281 13.7111C20.7381 13.3613 21.0371 12.9025 21.1893 12.3896C21.42 11.7036 21.5612 10.9915 21.6095 10.2705C21.6685 9.43855 21.731 8.46628 21.731 7.37419H24.2034V14.1219H26.1028V5.84731H20.026C20.026 7.13112 20.0087 8.21636 19.9809 9.07566ZM17.0189 5.86443H9.78577V14.139H11.6713V7.36049H15.0952V14.139H17.0015L17.0189 5.86443ZM43.9998 13.9884C44.5814 14.1976 45.197 14.2997 45.8159 14.2896C46.3584 14.3076 46.9005 14.2441 47.4237 14.1013C47.7957 13.9861 48.1516 13.8251 48.4828 13.622V12.1123C48.1351 12.3075 47.7657 12.4626 47.382 12.5745C46.9134 12.7055 46.428 12.7689 45.9409 12.7628C44.2209 12.7628 43.3609 11.843 43.3609 10.0034C43.3609 8.16387 44.2394 7.24181 45.9965 7.23725C46.4659 7.22895 46.9339 7.28893 47.3855 7.41527C47.7639 7.5273 48.124 7.69216 48.455 7.90483V6.38138C48.1398 6.16039 47.787 5.99691 47.4133 5.89866C46.9088 5.76882 46.3894 5.70324 45.868 5.70353C45.2326 5.70049 44.6012 5.80232 43.9998 6.00479C43.4767 6.17855 43.0013 6.46911 42.6109 6.85382C42.2194 7.23098 41.916 7.68779 41.7219 8.1924C41.3053 9.36246 41.3053 10.6375 41.7219 11.8076C41.9159 12.3132 42.2193 12.7712 42.6109 13.1496C43.0034 13.5293 43.4784 13.8162 43.9998 13.9884V13.9884ZM35.7701 12.7696C36.0565 12.7689 36.3398 12.7119 36.6035 12.6018C36.8717 12.4882 37.1094 12.3148 37.298 12.0952C37.6907 11.4674 37.8987 10.7444 37.8987 10.0068C37.8987 9.2693 37.6907 8.54627 37.298 7.91852C37.0045 7.59809 36.616 7.37701 36.1878 7.28667C35.7595 7.19632 35.3133 7.24132 34.9124 7.41527C34.6467 7.52989 34.4097 7.70049 34.2179 7.9151C33.8317 8.54573 33.6275 9.26843 33.6275 10.0051C33.6275 10.7418 33.8317 11.4645 34.2179 12.0952C34.4097 12.3098 34.6467 12.4804 34.9124 12.595C35.1744 12.7104 35.4589 12.7677 35.7458 12.7628V12.7628L35.7701 12.7696ZM57.7334 10.0034C57.7334 5.83362 57.7334 3.74872 56.6639 2.29031C56.3112 1.8173 55.8845 1.40246 55.3999 1.06128C53.9276 0 51.8337 0 47.6216 0L10.6226 0C6.41749 0 4.32013 0 2.84781 1.06128C2.36866 1.40419 1.94687 1.81888 1.59773 2.29031C0.535156 3.76583 0.535156 5.81992 0.535156 10.0034C0.535156 14.1869 0.535156 16.2547 1.59773 17.7131C1.9477 18.1848 2.36935 18.6005 2.84781 18.9456C4.32013 20 6.41749 20 10.6226 20H47.6216C51.8337 20 53.9276 20 55.4069 18.9456C55.8821 18.6003 56.3003 18.1845 56.6465 17.7131C57.716 16.2547 57.716 14.1664 57.716 10.0034H57.7334Z"
          fill="white"
        ></path>
        <path
          fill-rule="evenodd"
          clip-rule="evenodd"
          d="M136.358 1.79146L129.441 18.7925H123.173L125.069 14.1263C117.902 15.9099 111.575 11.5724 114.256 5.18414L115.645 1.78119H121.895L120.336 5.46828C118.02 10.7199 126.549 10.5042 128.132 6.61515L130.122 1.79146H136.358ZM96.9215 15.0267C98.2112 15.0303 99.4862 14.7565 100.657 14.2244C101.829 13.6922 102.868 12.9145 103.703 11.9455C101.714 11.9455 95.0013 12.3221 95.1298 14.0989C95.168 14.6843 95.6437 15.0267 96.9215 15.0267ZM113.242 13.9996L111.297 18.7925C108.992 18.8918 105.481 19.015 104.259 17.1698C100.217 19.9085 87.0563 20.7096 88.8411 13.6162C90.1503 8.39536 99.3036 8.1386 105.113 7.71751C105.776 4.86575 98.4494 4.86233 97.4945 7.54634L91.4212 7.70382C92.765 3.04446 97.6785 1.36353 102.297 1.33271C111.631 1.2574 111.867 7.42652 109.721 12.5617C108.974 14.342 111.815 13.9996 113.239 13.9996H113.242ZM77.9272 13.9996L80.9517 6.56038H75.2465C73.8818 9.37448 71.9442 12.7432 69.6906 13.9996H77.9272ZM142.751 15.0267C144.041 15.0303 145.316 14.7565 146.487 14.2244C147.658 13.6922 148.698 12.9145 149.533 11.9455C147.543 11.9455 140.834 12.3221 140.959 14.0989C140.997 14.6843 141.473 15.0267 142.751 15.0267ZM159.071 13.9996L157.127 18.7925C154.821 18.8918 151.311 19.015 150.088 17.1698C146.05 19.9085 132.886 20.7096 134.671 13.6162C135.983 8.39536 145.137 8.1386 150.946 7.71751C151.606 4.86575 144.282 4.86233 143.324 7.54634L137.251 7.70382C138.594 3.04446 143.508 1.36353 148.126 1.33271C157.46 1.2574 157.696 7.42652 155.55 12.5617C154.811 14.342 157.644 13.9996 159.068 13.9996H159.071ZM84.0665 13.9996H87.3237L85.3444 18.7925H57.6689L59.6135 13.9996C65.6174 13.9517 69.8989 7.34093 71.34 1.79146H89.0495L84.0665 13.9996Z"
          fill="white"
        ></path>
      </svg>
      <Button component={Link} to="/">
        Пик
      </Button>
      <Button component={Link} to="/settings/list">
        Настройки
      </Button>
    </div>
  );
};

export default TopBar;