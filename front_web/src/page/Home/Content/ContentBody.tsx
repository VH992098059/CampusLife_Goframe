import React from 'react';
import { Flex} from 'antd';
import CarouselHeader from "./CarouselHeader/CarouselHeader.tsx";
import FlexCenter from "./FlexCenter/FlexCenter.tsx";
import "./ContentBody.scss";


const ContentBody: React.FC = () => {
    return (
        <Flex gap="middle" vertical id={"carousel"}>
            <CarouselHeader/>
            <FlexCenter/>
        </Flex>
    );
};

export default ContentBody;
